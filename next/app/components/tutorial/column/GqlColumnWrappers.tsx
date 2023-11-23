import { nonNullArray } from "@/libs/nonNullArray";
import styles from "./GqlColumnWrappers.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { ColumnCarousel } from "./ColumnCarousel";
import { GqlColumnWrapper } from "./GqlColumnWrapper";

const fragmentDefinition = graphql(`
  fragment GqlColumnWrappers on Page {
    columns {
      columnName
      ...GqlColumnWrapper
    }
    focusColumn
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlColumnWrappers(props: Props): JSX.Element {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.columns) {
    return <div className={styles.empty}></div>;
  }

  const columns = nonNullArray(fragment.columns);
  const columnNames = columns.map((c) => c.columnName);
  const focusColumn = fragment.focusColumn ? fragment.focusColumn : "";

  return (
    <ColumnCarousel columnNames={columnNames} defaultFocusColumn={focusColumn}>
      <div className={styles.component}>
        {columns.map((c) => (
          <div key={c.columnName} className={styles.column}>
            <GqlColumnWrapper fragment={c} />
          </div>
        ))}
      </div>
    </ColumnCarousel>
  );
}
