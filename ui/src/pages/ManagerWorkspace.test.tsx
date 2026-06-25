import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ManagerWorkspace } from "./ManagerWorkspace";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useBalance: () => ({ data: undefined }),
  useReadContract: () => ({ data: undefined, isLoading: false, refetch: vi.fn() }),
  useReadContracts: () => ({ data: undefined, isLoading: false, refetch: vi.fn() }),
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn(), txs: [], dismiss: vi.fn() }) }));

describe("ManagerWorkspace", () => {
  it("renders the manager tab bar", () => {
    render(
      <QueryClientProvider client={new QueryClient()}>
        <MemoryRouter initialEntries={["/manager"]}>
          <Routes>
            <Route path="manager" element={<ManagerWorkspace />} />
          </Routes>
        </MemoryRouter>
      </QueryClientProvider>,
    );
    expect(screen.getByRole("link", { name: /service registry/i })).toBeInTheDocument();
    expect(screen.getByRole("link", { name: /booking token/i })).toBeInTheDocument();
  });
});
