import { nonNullArray } from "@/libs/nonNullArray";
import styles from "./GqlColumnWrappers.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { ColumnCarousel } from "./ColumnCarousel";
import { GqlColumnWrapper } from "./GqlColumnWrapper";

const fragmentDefinition = graphql(`
  fragment GqlColumnWrappers on Page {
    focusColumn
    columns {
      columnName
      ...GqlColumnWrapper
    }
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
  const defaultFocusColumn = fragment.focusColumn
    ? fragment.focusColumn
    : undefined;

  return (
    <ColumnCarousel
      columnNames={columnNames}
      defaultFocusColumn={defaultFocusColumn}
    >
      {/* This <div> is needed as <Carousel> expects single-element children */}
      <div className={styles.component}>
        {columns.map((c) => (
          // TODO: is this <div> needed?
          <div key={c.columnName} className={styles.column}>
            <GqlColumnWrapper
              fragment={c}
              defaultFocusColumn={defaultFocusColumn}
            />
          </div>
        ))}
      </div>
    </ColumnCarousel>
  );
}
