"use client";

import { nonNullArray } from "@/libs/nonNullArray";
import { useSearchParams } from "next/navigation";
import { useEffect, useRef, useState } from "react";
import { ColumnWrapperComponent } from "./ColumnWrapperComponent";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment Carousel_Fragment on Page {
    columns {
      ...ColumnWrapperComponent_Fragment
      name
    }
    focusColumn
    step
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
  columnIndex: number;
}

type State = Static | Animating;

interface CarouselProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  step: string;
  selectColumn?: string;
  openFilePath?: string;
  skipAnimation?: boolean;
}

export const Carousel = (props: CarouselProps) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const searchParams = useSearchParams();
  const ref = useRef<HTMLDivElement>(null);
  const columnParam = searchParams.get("column");
  const [state, setState] = useState<State>({ kind: "Static", columnIndex: 0 });

  useEffect(() => {
    switch (state.kind) {
      case "Animating":
        if (ref) {
          ref.current?.scrollIntoView({ behavior: "smooth", block: "end" });
        }
        setState({ kind: "Static", columnIndex: state.toIndex });
        break;
      case "Static":
        if (fragment?.columns && fragment.columns.length > 0) {
          const columns = nonNullArray(fragment.columns);

          // function's early return makes this logic clean - this is still cleaner than that of non-funcion :(
          const findIndex = (): number => {
            // 1st priority = 'column' query param;
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
            console.log(
              `Carousel animating from = ${state.columnIndex} , to = ${index}`
            );
            setState({
              kind: "Animating",
              fromIndex: state.columnIndex,
              toIndex: index,
              columnIndex: index,
            });
          }
        }
        break;
    }
  }, [fragment.columns, fragment.focusColumn, state, columnParam, props.step]);

  if (!fragment?.columns || fragment.columns.length === 0) {
    return <div></div>;
  }
  const columns = nonNullArray(fragment.columns);

  return (
    <div className={styles.carousel}>
      {columns.map((column, index) => (
        <div
          ref={index === state.columnIndex ? ref : null}
          key={column.name}
          className={styles.carouselElement}
        >
          <ColumnWrapperComponent
            fragment={column}
            step={props.step}
            skipAnimation={props.skipAnimation}
          />
        </div>
      ))}
    </div>
  );
};
