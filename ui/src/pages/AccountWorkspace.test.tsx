import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { AccountWorkspace } from "./AccountWorkspace";

vi.mock("wagmi", () => ({
  useChainId: () => 84532,
  useBalance: () => ({ data: { formatted: "1.0", symbol: "ETH" } }),
  useAccount: () => ({ address: undefined, chainId: 84532 }),
  useReadContract: () => ({ data: undefined, isLoading: false }),
  useReadContracts: () => ({ data: undefined, isLoading: false }),
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
  usePublicClient: () => undefined,
}));

vi.mock("../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));

describe("AccountWorkspace", () => {
  const addr = "0x1111111111111111111111111111111111111111";

  it("renders the tab bar and the account summary with the full address", () => {
    render(
      <QueryClientProvider client={new QueryClient()}>
        <MemoryRouter initialEntries={[`/account/${addr}`]}>
          <Routes>
            <Route path="account/:address" element={<AccountWorkspace />} />
          </Routes>
        </MemoryRouter>
      </QueryClientProvider>,
    );
    expect(screen.getByRole("link", { name: /bots/i })).toBeInTheDocument();
    // Full address appears in the left-pane summary.
    expect(screen.getAllByText(addr).length).toBeGreaterThan(0);
  });
});
