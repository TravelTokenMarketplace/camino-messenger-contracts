import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BookingTokenTab } from "./BookingTokenTab";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useReadContract: ({ functionName }: { functionName: string }) => {
    const map: Record<string, unknown> = {
      name: "BookingToken",
      symbol: "TRIP",
      version: [1n, 0n, 0n],
      getManagerAddress: "0x2222222222222222222222222222222222222222",
      getMinExpirationTimestampDiff: 60n,
    };
    return { data: map[functionName], isLoading: false, refetch: vi.fn() };
  },
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn() }) }));

function wrap(ui: React.ReactNode) {
  return render(<QueryClientProvider client={new QueryClient()}>{ui}</QueryClientProvider>);
}

describe("BookingTokenTab", () => {
  it("shows token info and settings", () => {
    wrap(<BookingTokenTab />);
    expect(screen.getByText("BookingToken")).toBeInTheDocument();
    expect(screen.getByText("TRIP")).toBeInTheDocument();
    expect(screen.getAllByText(/min expiration/i).length).toBeGreaterThan(0);
  });
});
