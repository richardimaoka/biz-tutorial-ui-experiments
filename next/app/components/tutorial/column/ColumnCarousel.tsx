"use client";

import { useSearchParams } from "next/navigation";
import { ReactNode } from "react";
import { Carousel } from "../../carousel/Carousel";
import { columnWidthPx } from "../definitions";

interface Props {
  columnNames: string[];
  children: ReactNode;
}

export function ColumnCarousel(props: Props) {
  const searchParams = useSearchParams();

  const currentColumn = searchParams.get("column");
  const foundIndex = props.columnNames.findIndex((c) => c === currentColumn);
  const currentIndex = foundIndex > -1 ? foundIndex : 0;

  return (
    <Carousel currentIndex={currentIndex} columnWidth={columnWidthPx}>
      {props.children}
    </Carousel>
  );
}
