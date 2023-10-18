import { ButtonToInitialStep } from "./ButtonToInitialStep";
import { ColumnTabs } from "./ColumnTabs";
import { ColumnName } from "../definitions";
import styles from "./TutorialHeader.module.css";

interface Props {
  selectTab: ColumnName;
  tabs: {
    name: ColumnName;
    href: string;
  }[];
}

export async function TutorialHeader(props: Props) {
  return (
    <div className={styles.component}>
      <ColumnTabs tabs={props.tabs} selectTab={props.selectTab} />
      <ButtonToInitialStep href="" />
    </div>
  );
}
