import { describe, expect, it } from "vitest";
import { render, screen } from "@testing-library/react";
import { RoleGate } from "./RoleGate";

describe("RoleGate", () => {
  it("renders children when role present", () => {
    render(
      <RoleGate hasRole roleName="BOT_ADMIN_ROLE">
        <button>Add bot</button>
      </RoleGate>,
    );
    expect(screen.getByRole("button", { name: /add bot/i })).toBeInTheDocument();
  });

  it("explains the missing role otherwise", () => {
    render(
      <RoleGate hasRole={false} roleName="BOT_ADMIN_ROLE">
        <button>Add bot</button>
      </RoleGate>,
    );
    expect(screen.queryByRole("button", { name: /add bot/i })).toBeNull();
    expect(screen.getByText(/BOT_ADMIN_ROLE/)).toBeInTheDocument();
  });

  it("names the gated action when provided", () => {
    render(
      <RoleGate hasRole={false} roleName="BOT_ADMIN_ROLE" action="Add bot">
        <span />
      </RoleGate>,
    );
    expect(screen.getByText(/can't add bot/i)).toBeInTheDocument();
  });
});
