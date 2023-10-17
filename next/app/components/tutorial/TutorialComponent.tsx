import { Carousel } from "./Carousel";
import styles from "./Tutorial.module.css";
import { Columns } from "./Columns";
import { TutorialHeader } from "./TutorialHeader";
import { ColumnName, TutorialColumnProps, columnWidthPx } from "./definitions";

interface Props {
  selectColumn: ColumnName;
  children: React.ReactNode;
  tabs: {
    name: ColumnName;
    href: string;
  }[];
  columns: TutorialColumnProps[];
  // skipAnimation?: boolean;
}

export async function TutorialComponent(props: Props) {
  return (
    <div className={styles.component}>
      {/* header part */}
      <div className={styles.header}>
        <TutorialHeader tabs={props.tabs} selectTab={props.selectColumn} />
      </div>
      {/* contents part */}
      <div className={styles.contents}>
        <Carousel currentIndex={8} columnWidth={columnWidthPx}>
          <Columns columns={props.columns} />
        </Carousel>
      </div>
    </div>
  );
}
