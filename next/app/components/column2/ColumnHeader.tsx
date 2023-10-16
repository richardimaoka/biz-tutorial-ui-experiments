import { ButtonToInitialStep } from "./ButtonToInitialStep";
import { ColumnTabs } from "./ColumnTabs";
import { TabName } from "./tabTypes";
import styles from "./ColumnHeader.module.css";

interface Props {
  selectTab: TabName;
  tabs: {
    name: TabName;
    href: string;
  }[];
}

export function ColumnHeader(props: Props) {
  return (
    <div className={styles.component}>
      <ColumnTabs tabs={props.tabs} selectTab={props.selectTab} />
      <ButtonToInitialStep href="" />
    </div>
  );
}
