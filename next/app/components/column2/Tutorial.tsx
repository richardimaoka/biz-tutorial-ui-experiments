import { Carousel } from "./Carousel";
import styles from "./Tutorial.module.css";
import { ColumnContents } from "./ColumnContents";
import { ColumnHeader } from "./ColumnHeader";
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
    //  Column consists of two parts,ColumnHeader and Carousel
    <div className={styles.component}>
      <div className={styles.headerHeight}>
        <ColumnHeader tabs={props.tabs} selectTab={props.selectColumn} />
      </div>
      <div className={styles.carouselHeight}>
        <Carousel currentIndex={0}>
          <div style={{ display: "flex", height: "100%" }}>
            <ColumnContents>
              <div>0</div>
            </ColumnContents>
            <ColumnContents>
              <div>1</div>
            </ColumnContents>
            <ColumnContents>
              <div>2</div>
            </ColumnContents>
            <ColumnContents>
              <div>3</div>
            </ColumnContents>
          </div>
        </Carousel>
      </div>
    </div>
  );
}
