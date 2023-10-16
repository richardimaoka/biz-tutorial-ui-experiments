"use client";

import React from "react";
import styles from "./Carousel.module.css";
import { columnWidthPx } from "./definitions";

interface Props {
  children: React.ReactNode;
  currentIndex: number;
}

export function Carousel(props: Props) {
  // Carousel is a client component, having client-side state
  const [currentIndex, setCurrentIndex] = React.useState(0);

  React.useEffect(() => {
    setCurrentIndex(props.currentIndex);
  }, [props.currentIndex]);

  // Supported Pattern: Passing Server Components to Client Components as Props:
  //   https://nextjs.org/docs/app/building-your-application/rendering/composition-patterns#supported-pattern-passing-server-components-to-client-components-as-props
  return (
    <div className={styles.component}>
      <div
        // carousel slider
        style={{
          transition: "transform 0.3s ease-in-out",
          transform: `translate(-${columnWidthPx * props.currentIndex}px)`,
        }}
      >
        {props.children}
      </div>
    </div>
  );
}
