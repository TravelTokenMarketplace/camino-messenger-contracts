export interface ParsedService {
  full: string;
  /** Package segment, e.g. "accommodation", "seat_map". Empty if unknown. */
  pkg: string;
  /** Version segment, e.g. "v5", "v1alpha". Undefined if none detected. */
  version?: string;
  /** Service name, e.g. "AccommodationSearchService" (falls back to full). */
  name: string;
}

/**
 * Parses a CMP service identifier of the form
 * `cmp.services.<package>.<version>.<ServiceName>` into its parts. Tolerant of
 * names that don't match: the whole string is used as the name in that case.
 */
export function parseServiceName(full: string): ParsedService {
  const parts = full.split(".");
  if (parts.length < 3) return { full, pkg: "", name: full };

  const name = parts[parts.length - 1];
  const maybeVersion = parts[parts.length - 2];
  const isVersion = /^v\d+\w*$/i.test(maybeVersion);
  const version = isVersion ? maybeVersion : undefined;

  const end = isVersion ? parts.length - 2 : parts.length - 1;
  const start = parts[0] === "cmp" && parts[1] === "services" ? 2 : 0;
  const pkg = parts.slice(start, end).join(".");

  return { full, pkg, version, name };
}

/** Numeric value of a version string for sorting ("v5" -> 5, "v1alpha" -> 1). */
export function versionOrder(version?: string): number {
  if (!version) return -1;
  const m = version.match(/\d+/);
  return m ? Number(m[0]) : -1;
}

export interface ServiceGroup<T> {
  pkg: string;
  items: (T & { parsed: ParsedService })[];
}

/**
 * Groups service-bearing items by package and sorts by package, then version,
 * then name — so related services cluster together and the distinguishing parts
 * line up. Each returned item is augmented with its parsed form.
 */
export function groupServicesByPackage<T extends { name: string }>(items: T[]): ServiceGroup<T>[] {
  const enriched = items.map((it) => ({ ...it, parsed: parseServiceName(it.name) }));
  enriched.sort(
    (a, b) =>
      a.parsed.pkg.localeCompare(b.parsed.pkg) ||
      versionOrder(a.parsed.version) - versionOrder(b.parsed.version) ||
      a.parsed.name.localeCompare(b.parsed.name),
  );
  const groups: ServiceGroup<T>[] = [];
  for (const s of enriched) {
    const key = s.parsed.pkg || "other";
    let g = groups.find((x) => x.pkg === key);
    if (!g) {
      g = { pkg: key, items: [] };
      groups.push(g);
    }
    g.items.push(s);
  }
  return groups;
}
