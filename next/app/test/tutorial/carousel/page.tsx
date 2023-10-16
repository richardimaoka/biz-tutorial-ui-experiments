import { Carousel } from "@/app/components/tutorial/Carousel";
import { columnWidthPx } from "@/app/components/tutorial/definitions";

export default async function Page() {
  const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
  return (
    <Carousel currentIndex={1}>
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
              width: `${columnWidthPx}px`,
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
  );
}
