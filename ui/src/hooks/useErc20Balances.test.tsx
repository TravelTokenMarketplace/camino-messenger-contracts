import { afterEach, describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";
import type { Address } from "viem";
import type { TokenMeta } from "./useTokenMetadata";

let mockSupported: unknown;
let mockBalances: { data: unknown; isLoading: boolean };
let mockMeta: Map<string, TokenMeta>;
const warnSpy = vi.fn();

vi.mock("wagmi", () => ({
  useReadContract: () => ({ data: mockSupported, isLoading: false }),
  useReadContracts: () => mockBalances,
}));
vi.mock("./useTokenMetadata", () => ({ useTokenMetadata: () => ({ meta: mockMeta, isLoading: false }) }));
vi.mock("./useActiveContracts", () => ({ useActiveContracts: () => ({ chainId: 84532, cmAccountAbi: [] }) }));
vi.mock("../config/tokens", () => ({ EXTRA_TOKENS: { 84532: ["0xAAaA000000000000000000000000000000000001"] } }));

import { useErc20Balances } from "./useErc20Balances";

const account = "0x1111111111111111111111111111111111111111" as const;
const A = "0xAAaA000000000000000000000000000000000001";
const metaOf = (over: Partial<TokenMeta>): Map<string, TokenMeta> =>
  new Map([[A.toLowerCase(), { address: A as Address, symbol: "USDC", name: "USD Coin", decimals: 6, ...over }]]);
const bal = (b: bigint) => ({ data: [{ status: "success", result: b }], isLoading: false });

afterEach(() => {
  mockSupported = [];
  mockBalances = { data: undefined, isLoading: false };
  mockMeta = metaOf({});
  warnSpy.mockClear();
  vi.unstubAllGlobals();
});

describe("useErc20Balances", () => {
  it("returns formatted balances, name and isZero flags", () => {
    mockSupported = [];
    mockMeta = metaOf({});
    mockBalances = bal(1500000n);
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(1);
    expect(result.current.tokens[0]).toMatchObject({
      symbol: "USDC",
      name: "USD Coin",
      decimals: 6,
      formatted: "1.5",
      isZero: false,
    });
  });

  it("flags zero balances", () => {
    mockBalances = bal(0n);
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens[0].isZero).toBe(true);
  });

  it("dedupes a config token that is also a supported token", () => {
    mockSupported = ["0xaaaa000000000000000000000000000000000001"];
    mockBalances = bal(1n);
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(1);
  });

  it("drops a token whose balanceOf read fails and warns in dev", () => {
    vi.stubGlobal("console", { ...console, warn: warnSpy });
    mockBalances = { data: [{ status: "failure", error: new Error("x") }], isLoading: false };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(0);
    expect(warnSpy).toHaveBeenCalled();
  });

  it("falls back to a short address when symbol metadata is missing", () => {
    mockMeta = metaOf({ symbol: undefined });
    mockBalances = bal(5000000n);
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens[0].symbol).toContain("…");
  });

  it("drops a token when decimals metadata is missing (avoids misformatting)", () => {
    vi.stubGlobal("console", { ...console, warn: warnSpy });
    mockMeta = metaOf({ decimals: undefined });
    mockBalances = bal(1500000n);
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(0);
    expect(warnSpy).toHaveBeenCalled();
  });
});
