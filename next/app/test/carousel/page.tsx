"use client";
import { ArrowLeftIcon } from "@/app/components/icons/ArrowLeftIcon";
import { ArrowRightIcon } from "@/app/components/icons/ArrowRightIcon";
import { Carousel } from "@/app/components/carousel/Carousel";
import React from "react";

export default function Page() {
  const arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
  const [columnIndex, setColumnIndex] = React.useState(0);

  return (
    <div style={{ width: "600px" }}>
      <Carousel currentIndex={columnIndex}>
        <div
          style={{
            display: "flex",
          }}
        >
          {arr.map((n) => (
            <div
              key={n}
              style={{
                backgroundColor: "white",
                width: "600px",
                height: "200px",
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                fontSize: "30px",
                flexShrink: 0,
              }}
            >
              {n}
            </div>
          ))}
        </div>
      </Carousel>
      <div
        style={{
          padding: "10px",
          display: "flex",
          justifyContent: "space-between",
        }}
      >
        <button
          onClick={() => {
            if (columnIndex > 0) {
              setColumnIndex(columnIndex - 1);
            }
          }}
        >
          <ArrowLeftIcon />
        </button>
        <button
          onClick={() => {
            if (columnIndex < arr.length - 1) {
              setColumnIndex(columnIndex + 1);
            }
          }}
        >
          <ArrowRightIcon />
        </button>
      </div>
    </div>
  );
}
