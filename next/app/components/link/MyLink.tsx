"use client";

import Link from "next/link";
import {
  ReadonlyURLSearchParams,
  usePathname,
  useSearchParams,
} from "next/navigation";
import { ReactNode } from "react";

interface Props {
  path?: string;
  searchParams?: Record<string, string>;
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
export function MyLink(props: Props) {
  // need to use router.replace instaed of <Link> not to mess up the browser history
  const pathname = usePathname();
  const searchParams = useSearchParams();

  const newPath = props.path ? props.path : pathname;
  const newSearchParams = replaceSearchParams(searchParams, props.searchParams);

  const href = newPath + "?" + newSearchParams.toString();

  // Prerequisites for props
  if (!props.path && !props.searchParams) {
    throw new Error("At least one of path and searchParams has to be passed");
  } else if (props.path && !props.path?.startsWith("/")) {
    throw new Error("path needs to start with '/'");
  }

  return (
    <Link href={href} replace>
      {props.children}
    </Link>
  );
}
