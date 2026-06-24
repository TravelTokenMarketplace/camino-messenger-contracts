import { afterEach, describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { RolesPanel } from "./RolesPanel";

// Mutable mock state so individual tests can simulate a connected admin wallet.
// (vitest only allows factory-referenced vars prefixed with `mock`.)
let mockAddress: string | undefined;
let mockReadData: unknown;

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: mockAddress }),
  useReadContract: () => ({ data: mockReadData, isLoading: false, refetch: vi.fn() }),
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

afterEach(() => {
  mockAddress = undefined;
  mockReadData = undefined;
});

describe("RolesPanel", () => {
  it("renders a row per role (enumerable)", () => {
    wrap(<RolesPanel address={addr} abi={[]} roles={["DEFAULT_ADMIN_ROLE", "PAUSER_ROLE"]} enumerable />);
    expect(screen.getByText("Admin")).toBeInTheDocument();
    expect(screen.getByText("Pauser")).toBeInTheDocument();
  });

  it("notes members are not listable in non-enumerable mode", () => {
    wrap(<RolesPanel address={addr} abi={[]} roles={["DEFAULT_ADMIN_ROLE"]} enumerable={false} />);
    fireEvent.click(screen.getByRole("button", { name: /admin/i }));
    expect(screen.getByText(/cannot list/i)).toBeInTheDocument();
  });

  it("renders grant/revoke-by-address inputs for an admin in non-enumerable mode", () => {
    // Connected wallet that holds the admin role: useHasRole reads truthy.
    mockAddress = addr;
    mockReadData = true;
    wrap(<RolesPanel address={addr} abi={[]} roles={["DEFAULT_ADMIN_ROLE"]} enumerable={false} />);
    fireEvent.click(screen.getByRole("button", { name: /admin/i }));
    expect(screen.getByPlaceholderText("Address 0x…")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Address 0x… to revoke")).toBeInTheDocument();
  });
});
