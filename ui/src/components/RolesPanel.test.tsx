import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { RolesPanel } from "./RolesPanel";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useReadContract: () => ({ data: undefined, isLoading: false, refetch: vi.fn() }),
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn() }) }));

const addr = "0x1111111111111111111111111111111111111111" as const;

function wrap(ui: React.ReactNode) {
  return render(<QueryClientProvider client={new QueryClient()}>{ui}</QueryClientProvider>);
}

describe("RolesPanel", () => {
  it("renders a row per role (enumerable)", () => {
    wrap(<RolesPanel address={addr} abi={[]} roles={["DEFAULT_ADMIN_ROLE", "PAUSER_ROLE"]} enumerable />);
    expect(screen.getByText("Admin")).toBeInTheDocument();
    expect(screen.getByText("Pauser")).toBeInTheDocument();
  });

  it("notes members are not listable in non-enumerable mode", () => {
    wrap(<RolesPanel address={addr} abi={[]} roles={["DEFAULT_ADMIN_ROLE"]} enumerable={false} />);
    expect(screen.getByText(/cannot list/i)).toBeInTheDocument();
  });
});
