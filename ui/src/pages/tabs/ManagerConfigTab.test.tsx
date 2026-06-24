import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ManagerConfigTab } from "./ManagerConfigTab";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useReadContract: ({ functionName }: { functionName: string }) => ({
    data: functionName === "paused" ? false : undefined,
    isLoading: false,
    refetch: vi.fn(),
  }),
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn() }) }));

function wrap(ui: React.ReactNode) {
  return render(<QueryClientProvider client={new QueryClient()}>{ui}</QueryClientProvider>);
}

describe("ManagerConfigTab", () => {
  it("shows pause status and config sections", () => {
    wrap(<ManagerConfigTab />);
    expect(screen.getAllByText(/account creation/i).length).toBeGreaterThan(0);
    expect(screen.getAllByText(/account implementation/i).length).toBeGreaterThan(0);
    expect(screen.getAllByText(/booking token address/i).length).toBeGreaterThan(0);
  });
});
