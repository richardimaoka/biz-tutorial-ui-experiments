import { Carousel } from "@/app/components/column2/Carousel";

export default async function Page() {
  const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
  return (
    <Carousel fromIndex={0} toIndex={0}>
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
              width: "400px",
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
