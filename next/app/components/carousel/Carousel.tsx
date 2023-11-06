import styles from "./Carousel.module.css";

interface Props {
  children: React.ReactNode;
  currentIndex: number;
  columnWidth: number;
}

export function Carousel(props: Props) {
  // `-0%` makes it a bit awkward, so special handling for 0%
  const translatePercentage =
    props.currentIndex === 0 ? 0 : -100 * props.currentIndex;

  // TODO: might throw on currentIndex > maxIndex, to render error route
  return (
    <div className={styles.component}>
      <div
        className={styles.carouselSlider}
        style={{
          // simple props can make client-side animations!!
          transition: "transform 0.3s ease-in-out",
          transform: `translate(${translatePercentage}%)`,
        }}
      >
        {/* props.children should be wider-than <Carousel> */}
        {/* props.children should be an array of <Column>'s */}
        {props.children /* children, for loose coupling */}
      </div>
    </div>
  );
}
