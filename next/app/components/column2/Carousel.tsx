"use client";

import React from "react";
import styles from "./Carousel.module.css";

interface Props {
  children: React.ReactNode;
  fromIndex: number;
  toIndex: number;
}

export function Carousel(props: Props) {
  // Carousel is a client component, having client-side state
  const [fromIndex, setFromIndex] = React.useState(0);
  const [toIndex, setToIndex] = React.useState(0);
  console.log("props.toIndex", props.toIndex);

  React.useEffect(() => {
    setFromIndex(props.fromIndex);
    setToIndex(props.toIndex);
  }, [props.fromIndex, props.toIndex]);

  // Supported Pattern: Passing Server Components to Client Components as Props:
  //   https://nextjs.org/docs/app/building-your-application/rendering/composition-patterns#supported-pattern-passing-server-components-to-client-components-as-props
  return (
    <div className={styles.component}>
      <div
        style={{
          transition: "all 0.3s ease-in-out",
          transform: `translate(${-400 * props.toIndex}px)`,
        }}
      >
        {props.children}
      </div>
    </div>
  );
}
