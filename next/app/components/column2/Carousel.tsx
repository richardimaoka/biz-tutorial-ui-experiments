import styles from "./Carousel.module.css";
import { columnWidthPx } from "./definitions";

interface Props {
  children: React.ReactNode;
  currentIndex: number;
}

export async function Carousel(props: Props) {
  // TODO: might throw on currentIndex > maxIndex, to render error route
  return (
    <div className={styles.component}>
      <div
        className={styles.carouselSlider}
        style={{
          // simple props can make client-side animations!!
          transition: "transform 0.3s ease-in-out",
          transform: `translate(-${columnWidthPx * props.currentIndex}px)`,
        }}
      >
        {/* props.children should be wider-than <Carousel> */}
        {/* props.children should be an array */}
        {props.children}
      </div>
    </div>
  );
}
