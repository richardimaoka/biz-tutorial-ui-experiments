import { ButtonToInitialStep } from "./ButtonToInitialStep";
import { ColumnTabProps } from "./ColumnTab";
import { ColumnTabs } from "./ColumnTabs";

interface Props {
  tabs: ColumnTabProps[];
}

export function ColumnHeader(props: Props) {
  return (
    <div>
      <ColumnTabs tabs={props.tabs} />
      <ButtonToInitialStep href="" />
    </div>
  );
}
