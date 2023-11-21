"use client";

import Link from "next/link";
import {
  ReadonlyURLSearchParams,
  usePathname,
  useSearchParams,
} from "next/navigation";
import { ReactNode } from "react";

interface Props {
  searchParams: Record<string, string>;
  children: ReactNode;
}

function replaceSearchParams(
  existing: ReadonlyURLSearchParams,
  newParams?: Record<string, string>
) {
  const params = new URLSearchParams(existing);

  for (const name in newParams) {
    params.set(name, newParams[name]);
  }

  return params;
}

// https://nextjs.org/docs/app/api-reference/functions/use-search-params#examples
export function LinkSearchParams(props: Props) {
  const pathname = usePathname();
  const searchParams = useSearchParams();

  const newSearchParams = replaceSearchParams(searchParams, props.searchParams);

  const href = pathname + "?" + newSearchParams.toString();

  return (
    <Link href={href} replace>
      {props.children}
    </Link>
  );
}
