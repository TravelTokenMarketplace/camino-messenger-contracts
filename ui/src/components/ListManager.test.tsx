import { describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { ListManager } from "./ListManager";

const base = {
  title: "Payment Tokens",
  isLoading: false,
  roleName: "DEFAULT_ADMIN_ROLE",
  addLabel: "Add token",
  addPlaceholder: "0x…",
  onRemove: vi.fn().mockResolvedValue("0xhash"),
};

describe("ListManager", () => {
  it("renders items and calls onAdd with the input value", async () => {
    const onAdd = vi.fn().mockResolvedValue("0xhash");
    render(<ListManager {...base} hasRole items={["0xAAA"]} onAdd={onAdd} />);
    expect(screen.getByText("0xAAA")).toBeInTheDocument();
    fireEvent.change(screen.getByPlaceholderText("0x…"), { target: { value: "0xBBB" } });
    fireEvent.click(screen.getByRole("button", { name: /add token/i }));
    await waitFor(() => expect(onAdd).toHaveBeenCalledWith("0xBBB"));
  });

  it("hides add/remove controls without the role", () => {
    render(<ListManager {...base} hasRole={false} items={["0xAAA"]} onAdd={vi.fn()} />);
    expect(screen.queryByRole("button", { name: /add token/i })).toBeNull();
    expect(screen.getAllByText(/DEFAULT_ADMIN_ROLE/).length).toBeGreaterThan(0);
  });
});
