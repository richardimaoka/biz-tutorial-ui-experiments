import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { ColumnHeader } from "./ColumnHeader";

import styles from "./style.module.css";
import { ColumnWrapperComponent } from "./ColumnWrapperComponent";
import { nonNullArray } from "@/libs/nonNullArray";

const fragmentDefinition = graphql(`
  fragment VisibleColumn_Fragment on Page {
    ...ColumnHeader_Fragment
    columns {
      ...ColumnWrapperComponent_Fragment
      name
    }
  }
`);

interface VisibleColumnProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  selectColumn?: string;
  openFilePath?: string;
  step?: string;
}

export const VisibleColumn = (props: VisibleColumnProps) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment?.columns) {
    return <div></div>;
  }
  const columns = nonNullArray(fragment.columns);

  const selectColumn = props.selectColumn
    ? decodeURI(props.selectColumn)
    : columns.length > 0 && columns[0].name
    ? columns[0].name
    : "";

  const visibleColumn = columns.find((column) => column.name === selectColumn);

  console.log(props.selectColumn);

  return (
    <div className={styles.visiblecolumn}>
      <ColumnHeader
        fragment={fragment}
        selectColumn={selectColumn}
        openFilePath={props.openFilePath}
        step={props.step}
      />
      <div className={styles.wrapper}>
        {/* above <div> + .wrapper style is necessary to control the height of visible column = 100svh */}
        {visibleColumn && <ColumnWrapperComponent fragment={visibleColumn} />}
      </div>
    </div>
  );
};
