import { describe, expect, it } from "vitest";
import { ACCOUNT_EVENTS, BOOKING_TOKEN_EVENTS, MANAGER_EVENTS, lookupEntry, toActivityEvent } from "./catalog";

const ACC = "0xaAaA000000000000000000000000000000000001";
const BUYER = "0xbBbB000000000000000000000000000000000002";

function render(source: Parameters<typeof lookupEntry>[0], name: string, args: Record<string, unknown>) {
  const entry = lookupEntry(source, name);
  if (!entry) throw new Error(`no catalog entry for ${source}:${name}`);
  return entry.render(args);
}

describe("catalog rendering", () => {
  it("renders manager account + service events", () => {
    expect(render("manager", "CMAccountCreated", { account: ACC })).toBe("CM Account 0xaAaA…0001 created");
    expect(render("manager", "ServiceRegistered", { serviceName: "cmp.x.v1.Foo" })).toBe(
      'Service "cmp.x.v1.Foo" registered',
    );
  });

  it("renders booking + cancellation events with the token id", () => {
    expect(render("bookingToken", "TokenBought", { tokenId: 42n, buyer: BUYER })).toBe(
      "Booking token #42 bought by 0xbBbB…0002",
    );
    expect(render("bookingToken", "TokenReservationExpired", { tokenId: 7n })).toBe(
      "Booking token #7 reservation expired",
    );
    expect(render("bookingToken", "CancellationFinalized", { tokenId: 9n })).toBe(
      "Cancellation finalized for booking token #9",
    );
  });

  it("renders account fund events with ether-formatted amounts", () => {
    expect(render("account", "Deposit", { sender: ACC, amount: 1_500000000000000000n })).toBe(
      "Deposit of 1.5 from 0xaAaA…0001",
    );
    expect(render("account", "OffChainPaymentSupportUpdated", { supportsOffChainPayment: true })).toBe(
      "Off-chain payment support enabled",
    );
  });

  it("omits the (hashed, indexed) service name for account service events", () => {
    expect(render("account", "ServiceAdded", { serviceName: "0xdeadbeef" })).toBe("Supported service added");
  });

  it("exposes non-overlapping event sets per source", () => {
    expect(MANAGER_EVENTS.length).toBe(3);
    expect(BOOKING_TOKEN_EVENTS.length).toBe(7);
    expect(ACCOUNT_EVENTS.length).toBeGreaterThan(15);
  });
});

describe("toActivityEvent", () => {
  const log = {
    eventName: "TokenBought",
    args: { tokenId: 5n, buyer: BUYER },
    blockNumber: 100n,
    logIndex: 3,
    transactionHash: "0xabc",
    address: "0xcontract",
  };

  it("maps a decoded log to a normalized event", () => {
    const ev = toActivityEvent(log as never, "bookingToken");
    expect(ev).toMatchObject({
      id: "0xabc#3",
      source: "bookingToken",
      category: "Bookings",
      blockNumber: 100n,
      logIndex: 3,
      sentence: "Booking token #5 bought by 0xbBbB…0002",
    });
  });

  it("returns undefined for an unknown event", () => {
    expect(toActivityEvent({ ...log, eventName: "Nope" } as never, "bookingToken")).toBeUndefined();
  });
});
