"use client";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { ColumnHeader } from "./ColumnHeader";

import styles from "./style.module.css";
import { ColumnWrapperComponent } from "./ColumnWrapperComponent";
import { nonNullArray } from "@/libs/nonNullArray";
import { ModalFrame } from "../modal/ModalFrame";
import { Navigation } from "../navigation/Navigation";
import { useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";

const fragmentDefinition = graphql(`
  fragment Carousel_Fragment on Page {
    columns {
      ...ColumnWrapperComponent_Fragment
      name
    }
    focusColumn
  }
`);

interface Static {
  kind: "Static";
  columnIndex: number;
}

interface Animating {
  kind: "Animating";
  fromIndex: number;
  toIndex: number;
}

type State = Static | Animating;

interface CarouselProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  step: string;
  selectColumn?: string;
  openFilePath?: string;
  skipAnimation?: boolean;
}

export const VisibleColumn = (props: CarouselProps) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const searchParams = useSearchParams();
  const [state, setState] = useState<State>({ kind: "Static", columnIndex: 0 });

  useEffect(() => {
    switch (state.kind) {
      case "Animating":
        break;
      case "Static":
        console.log("useEffect Static");
        if (fragment?.columns && fragment.columns.length > 0) {
          const columns = nonNullArray(fragment.columns);

          // function's early return makes this logic clean - this is still cleaner than that of non-funcion :(
          const findIndex = (): number => {
            // 1st priority = 'column' query param
            const columnParam = searchParams.get("column");
            if (columnParam) {
              const index = columns.findIndex(
                (col) => col.name === decodeURI(columnParam)
              );
              if (index > -1) {
                return index;
              }
            }

            // 2nd priority = 'focusColumn' field from server
            const focusColumn = fragment.focusColumn;
            if (focusColumn) {
              const index = columns.findIndex(
                (col) => col.name === decodeURI(focusColumn)
              );
              if (index > -1) {
                return index;
              }
            }

            return 0;
          };

          const index = findIndex();
          if (index !== state.columnIndex) {
            setState({
              kind: "Animating",
              fromIndex: state.columnIndex,
              toIndex: index,
            });
          }
        }
        break;
    }
  }, [fragment.columns, fragment.focusColumn, state, searchParams]);

  if (!fragment?.columns || fragment.columns.length === 0) {
    return <div></div>;
  }
  const columns = nonNullArray(fragment.columns);

  return (
    <div>
      {columns.map((column) => (
        <ColumnWrapperComponent
          key={column.name}
          fragment={column}
          step={props.step}
          skipAnimation={props.skipAnimation}
        />
      ))}
    </div>
  );
};
