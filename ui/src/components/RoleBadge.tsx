export function RoleBadge({ role }: { role: string }) {
  return (
    <span className="rounded bg-indigo-50 px-2 py-0.5 text-xs text-indigo-700 dark:bg-indigo-950 dark:text-indigo-300">
      {role}
    </span>
  );
}
