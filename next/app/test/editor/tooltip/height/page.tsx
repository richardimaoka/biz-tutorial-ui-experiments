"use client";

import { useEffect, useRef } from "react";

function Child() {
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    console.log("ref = ", ref);
  });

  return (
    <div
      ref={ref}
      style={{
        margin: "20px",
        width: "400px",
        height: "200px",
        backgroundColor: "white",
      }}
    >
      rect = {ref.current && ref.current.offsetHeight}
    </div>
  );
}

export default function Page() {
  return (
    <div>
      <Child />
    </div>
  );
}
