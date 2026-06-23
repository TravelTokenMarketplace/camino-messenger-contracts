import { describe, expect, it } from "vitest";
import { wagmiConfig } from "./wagmi";
import { ENABLED_CHAINS } from "../config/chains";

describe("wagmiConfig", () => {
  it("registers a transport for every enabled chain", () => {
    for (const c of ENABLED_CHAINS) {
      expect(wagmiConfig.chains.some((wc) => wc.id === c.id)).toBe(true);
    }
  });

  it("does not register the disabled Columbus chain", () => {
    expect(wagmiConfig.chains.some((wc) => wc.id === 501)).toBe(false);
  });
});
