import { describe, expect, it } from "vitest";
import { parseServiceName, versionOrder } from "./serviceName";

describe("parseServiceName", () => {
  it("parses a standard service identifier", () => {
    expect(parseServiceName("cmp.services.accommodation.v5.AccommodationSearchService")).toEqual({
      full: "cmp.services.accommodation.v5.AccommodationSearchService",
      pkg: "accommodation",
      version: "v5",
      name: "AccommodationSearchService",
    });
  });

  it("handles packages with underscores", () => {
    const p = parseServiceName("cmp.services.seat_map.v4.SeatMapService");
    expect(p.pkg).toBe("seat_map");
    expect(p.version).toBe("v4");
    expect(p.name).toBe("SeatMapService");
  });

  it("recognises alpha versions", () => {
    expect(parseServiceName("cmp.services.ping.v1alpha.PingService").version).toBe("v1alpha");
  });

  it("falls back gracefully for non-conforming names", () => {
    expect(parseServiceName("weird")).toEqual({ full: "weird", pkg: "", name: "weird" });
  });

  it("orders versions numerically", () => {
    expect(versionOrder("v5")).toBe(5);
    expect(versionOrder("v1alpha")).toBe(1);
    expect(versionOrder(undefined)).toBe(-1);
  });
});
