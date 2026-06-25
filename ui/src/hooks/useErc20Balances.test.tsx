import { afterEach, describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";

let mockSupported: unknown;
let mockMulticall: { data: unknown; isLoading: boolean };
const warnSpy = vi.fn();

vi.mock("wagmi", () => ({
  useReadContract: () => ({ data: mockSupported, isLoading: false }),
  useReadContracts: () => mockMulticall,
}));
vi.mock("../hooks/useActiveContracts", () => ({
  useActiveContracts: () => ({ chainId: 84532, cmAccountAbi: [] }),
}));
vi.mock("../config/tokens", () => ({ EXTRA_TOKENS: { 84532: ["0xAAaA000000000000000000000000000000000001"] } }));

import { useErc20Balances } from "./useErc20Balances";

const account = "0x1111111111111111111111111111111111111111" as const;
// 3 calls per token: symbol, decimals, balanceOf
const ok = (s: string, d: number, b: bigint) => [
  { status: "success", result: s },
  { status: "success", result: d },
  { status: "success", result: b },
];

afterEach(() => {
  mockSupported = undefined;
  mockMulticall = { data: undefined, isLoading: false };
  warnSpy.mockClear();
  vi.unstubAllGlobals();
});

describe("useErc20Balances", () => {
  it("returns formatted balances and isZero flags", () => {
    mockSupported = [];
    mockMulticall = { data: ok("USDC", 6, 1500000n), isLoading: false };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(1);
    expect(result.current.tokens[0]).toMatchObject({ symbol: "USDC", decimals: 6, formatted: "1.5", isZero: false });
  });

  it("flags zero balances", () => {
    mockSupported = [];
    mockMulticall = { data: ok("USDC", 6, 0n), isLoading: false };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens[0].isZero).toBe(true);
  });

  it("dedupes a config token that is also a supported token", () => {
    mockSupported = ["0xaaaa000000000000000000000000000000000001"]; // same as config, different case
    mockMulticall = { data: ok("USDC", 6, 1n), isLoading: false };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(1);
  });

  it("drops a non-ERC20 (balanceOf failure) and warns in dev", () => {
    vi.stubGlobal("console", { ...console, warn: warnSpy });
    mockSupported = [];
    mockMulticall = {
      data: [
        { status: "failure", error: new Error("x") },
        { status: "failure", error: new Error("x") },
        { status: "failure", error: new Error("x") },
      ],
      isLoading: false,
    };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(0);
    expect(warnSpy).toHaveBeenCalled();
  });

  it("keeps token with symbol fallback when only symbol fails", () => {
    mockSupported = [];
    mockMulticall = {
      data: [
        { status: "failure", error: new Error("x") }, // symbol
        { status: "success", result: 18 }, // decimals
        { status: "success", result: 5000000000000000000n }, // balanceOf
      ],
      isLoading: false,
    };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(1);
    expect(result.current.tokens[0]).toMatchObject({ decimals: 18, formatted: "5" });
    // symbol falls back to the shortened address (contains the ellipsis)
    expect(result.current.tokens[0].symbol).toContain("…");
  });

  it("drops a token when decimals can't be read (avoids misformatting)", () => {
    vi.stubGlobal("console", { ...console, warn: warnSpy });
    mockSupported = [];
    mockMulticall = {
      data: [
        { status: "success", result: "USDC" }, // symbol
        { status: "failure", error: new Error("x") }, // decimals
        { status: "success", result: 1500000n }, // balanceOf
      ],
      isLoading: false,
    };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(0);
    expect(warnSpy).toHaveBeenCalled();
  });
});
