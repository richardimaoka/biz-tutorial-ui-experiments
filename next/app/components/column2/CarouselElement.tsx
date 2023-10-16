interface Props {
  name: string;
}

export async function CarouselElement(props: Props) {
  const data = await fetch("https://google.com");
  return (
    <div style={{ width: "100px", height: "100px", backgroundColor: "white" }}>
      {props.name}
    </div>
  );
}
