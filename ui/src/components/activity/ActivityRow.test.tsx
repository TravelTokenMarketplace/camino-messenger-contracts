import { describe, expect, it } from "vitest";
import { render, screen } from "@testing-library/react";
import { ActivityRow } from "./ActivityRow";
import { type ActivityEvent } from "../../lib/activity/types";

const baseEvent: ActivityEvent = {
  id: "0xabc#1",
  source: "bookingToken",
  category: "Bookings",
  contract: "0x0000000000000000000000000000000000000001",
  blockNumber: 123n,
  logIndex: 1,
  txHash: "0xabc",
  eventName: "TokenBought",
  args: {},
  sentence: "Booking token #5 bought by 0xbBbB…0002",
};

describe("ActivityRow", () => {
  it("renders the sentence and a tx explorer link when timestamped", () => {
    const event = { ...baseEvent, timestamp: Math.floor(Date.now() / 1000) - 120 };
    render(<ActivityRow event={event} explorerUrl="https://explorer.test" />);

    expect(screen.getByText("Booking token #5 bought by 0xbBbB…0002")).toBeInTheDocument();
    const link = screen.getByRole("link", { name: "View transaction" });
    expect(link).toHaveAttribute("href", "https://explorer.test/tx/0xabc");
    expect(screen.getByText("2m ago")).toBeInTheDocument();
  });

  it("falls back to the block number when no timestamp is resolved", () => {
    render(<ActivityRow event={baseEvent} />);
    expect(screen.getByText("block 123")).toBeInTheDocument();
  });
});
