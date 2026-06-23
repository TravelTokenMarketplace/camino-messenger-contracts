import { describe, expect, it } from "vitest";
import { resolveAddresses } from "./sync-contracts";

describe("resolveAddresses", () => {
  it("resolves canonical proxy addresses per chain", () => {
    const out = resolveAddresses({
      84532: {
        "CaminoMessengerModule#ManagerProxy": "0xMANAGER",
        "CaminoMessengerModule#BookingTokenProxy": "0xBT",
        "CaminoMessengerModule#CMAccount": "0xIMPL",
        "CaminoMessengerModule#ManagerERC1967Proxy": "0xMANAGER",
      },
    });
    expect(out[84532]).toEqual({
      manager: "0xMANAGER",
      bookingToken: "0xBT",
      cmAccountImpl: "0xIMPL",
    });
  });

  it("ignores non-canonical historical modules (Columbus mess)", () => {
    const out = resolveAddresses({
      501: {
        "CaminoMessengerModule#ManagerProxy": "0xGOOD",
        "CaminoMessengerModule#BookingTokenProxy": "0xGOODBT",
        "CaminoMessengerModule#CMAccount": "0xGOODIMPL",
        "RefactorCancellationModule#ManagerProxy": "0xBAD",
        "ERC20ServiceFeeModule#CMAccount": "0xBAD2",
      },
    });
    expect(out[501].manager).toBe("0xGOOD");
    expect(out[501].cmAccountImpl).toBe("0xGOODIMPL");
  });

  it("omits chains missing required keys", () => {
    const out = resolveAddresses({ 8453: { "Other#Thing": "0x0" } });
    expect(out[8453]).toBeUndefined();
  });
});
