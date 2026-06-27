import { afterEach, describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { AccountSummary } from "./AccountSummary";
import type { TokenBalance } from "../hooks/useErc20Balances";

let mockNative: { formatted: string; symbol: string; value: bigint } | undefined;
let mockTokens: TokenBalance[];

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useBalance: () => ({ data: mockNative }),
}));
vi.mock("../hooks/useActiveContracts", () => ({
  useActiveContracts: () => ({ chainId: 84532 }),
}));
vi.mock("../hooks/useMyAccounts", () => ({ useAccountRolesFor: () => ({ roles: [], isLoading: false }) }));
vi.mock("../hooks/useErc20Balances", () => ({ useErc20Balances: () => ({ tokens: mockTokens, isLoading: false }) }));

const account = "0x1111111111111111111111111111111111111111" as const;
const tok = (over: Partial<TokenBalance>): TokenBalance => ({
  address: "0xAAaA000000000000000000000000000000000001",
  symbol: "USDC",
  decimals: 6,
  balance: 1500000n,
  formatted: "1.5",
  isZero: false,
  ...over,
});

function wrap(ui: React.ReactNode) {
  return render(<QueryClientProvider client={new QueryClient()}>{ui}</QueryClientProvider>);
}

afterEach(() => {
  mockNative = { formatted: "1.0", symbol: "ETH", value: 1000000000000000000n };
  mockTokens = [];
});

describe("AccountSummary ERC20 balances", () => {
  it("renders a row per token with formatted amount and symbol", () => {
    mockNative = { formatted: "1.0", symbol: "ETH", value: 1000000000000000000n };
    mockTokens = [tok({})];
    wrap(<AccountSummary account={account} />);
    // Tooltip renders a hidden duplicate of the amount; assert the visible (non-tooltip) span.
    expect(screen.getByText(/1\.5 USDC/, { selector: 'span:not([role="tooltip"])' })).toBeInTheDocument();
  });

  it("warns when an ERC20 balance is zero", () => {
    mockNative = { formatted: "1.0", symbol: "ETH", value: 1000000000000000000n };
    mockTokens = [tok({ balance: 0n, formatted: "0", isZero: true })];
    wrap(<AccountSummary account={account} />);
    expect(screen.getByText(/can't buy booking tokens paid in USDC/i)).toBeInTheDocument();
  });

  it("warns when native balance is zero", () => {
    mockNative = { formatted: "0", symbol: "ETH", value: 0n };
    mockTokens = [];
    wrap(<AccountSummary account={account} />);
    expect(screen.getByText(/no .* for gas|can't pay gas/i)).toBeInTheDocument();
  });
});
