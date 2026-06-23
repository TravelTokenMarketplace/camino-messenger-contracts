import { describe, expect, it } from "vitest";
import { APP_CHAINS, ENABLED_CHAINS } from "./chains";

describe("chain config", () => {
  it("defines all four networks", () => {
    expect(APP_CHAINS.map((c) => c.id).sort()).toEqual([500, 501, 8453, 84532]);
  });

  it("marks Columbus (501) as disabled", () => {
    expect(APP_CHAINS.find((c) => c.id === 501)!.enabled).toBe(false);
  });

  it("excludes Columbus from enabled chains", () => {
    expect(ENABLED_CHAINS.some((c) => c.id === 501)).toBe(false);
  });

  it("only enables chains that have deployed addresses", () => {
    for (const c of ENABLED_CHAINS) {
      expect(c.enabled).toBe(true);
    }
  });
});
