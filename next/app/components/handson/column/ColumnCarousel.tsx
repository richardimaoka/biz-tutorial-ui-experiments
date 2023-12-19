"use client";

import { useSearchParams } from "next/navigation";
import { ReactNode } from "react";
import { Carousel } from "../../carousel/Carousel";

interface Props {
  columnNames: string[];
  children: ReactNode;
  defaultFocusColumn?: string;
}

export function ColumnCarousel(props: Props) {
  const searchParams = useSearchParams();

  const columnParam = searchParams.get("column");
  // prefer the query parameter, if no query parameter, then default focus from GraphQL
  const currentColumn = columnParam ? columnParam : props.defaultFocusColumn;

  const foundIndex = props.columnNames.findIndex((c) => c === currentColumn);
  const currentIndex = foundIndex > -1 ? foundIndex : 0;

  return <Carousel currentIndex={currentIndex}>{props.children}</Carousel>;
}
