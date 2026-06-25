import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ServiceRegistryTab } from "./ServiceRegistryTab";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useReadContract: ({ functionName }: { functionName: string }) => ({
    data:
      functionName === "getAllRegisteredServiceNames"
        ? ["cmp.services.accommodation.v1.AccommodationSearchService"]
        : undefined,
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

describe("ServiceRegistryTab", () => {
  it("lists registered services", () => {
    wrap(<ServiceRegistryTab />);
    expect(screen.getByText("AccommodationSearchService")).toBeInTheDocument();
  });
});
