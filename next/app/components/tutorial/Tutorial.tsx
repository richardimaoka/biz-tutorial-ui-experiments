import { Carousel } from "./Carousel";
import styles from "./Tutorial.module.css";
import { Columns } from "./Columns";
import { TutorialHeader } from "./TutorialHeader";
import { ColumnName } from "./definitions";

interface Props {
  selectColumn: ColumnName;
  children: React.ReactNode;
  tabs: {
    name: ColumnName;
    href: string;
  }[];
  // selectColumn?: string;
  // skipAnimation?: boolean;
}

export async function Tutorial(props: Props) {
  return (
    //  Column consists of two parts,TutorialHeader and Carousel
    <div className={styles.component}>
      <div className={styles.headerHeight}>
        <TutorialHeader tabs={props.tabs} selectTab={props.selectColumn} />
      </div>
      <div className={styles.carouselHeight}>
        <Carousel currentIndex={8}>
          <Columns />
        </Carousel>
      </div>
    </div>
  );
}
