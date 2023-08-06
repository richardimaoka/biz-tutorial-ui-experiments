import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./style.module.css";
import { nonNullArray } from "@/libs/nonNullArray";
import { ColumnTab } from "./ColumnTab";

const fragmentDefinition = graphql(`
  fragment ColumnTabs_Fragment on Page {
    columns {
      ...ColumnTab_Fragment
      name
    }
  }
`);

export interface ColumnTabProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  selectColumn: string;
  openFilePath?: string;
  step?: string;
}

export const ColumnTabs = (props: ColumnTabProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.columns) {
    return <div></div>;
  }
  const columns = nonNullArray(fragment.columns);

  return (
    <div className={styles.tabs}>
      {columns.map((col) => (
        <ColumnTab
          key={col.name}
          isSelected={col.name === props.selectColumn}
          openFilePath={props.openFilePath}
          step={props.step}
          fragment={col}
        />
      ))}
    </div>
  );
};
