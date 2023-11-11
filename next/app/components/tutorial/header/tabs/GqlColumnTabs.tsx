import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { nonNullArray } from "@/libs/nonNullArray";
import { GqlColumnTab } from "./GqlColumnTab";
import styles from "./GqlColumnTabs.module.css";

const fragmentDefinition = graphql(`
  fragment GqlColumnTabs on Page {
    columns {
      columnName
      ...GqlColumnTab
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlColumnTabs(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.columns) {
    return <></>;
  }

  const columns = nonNullArray(fragment.columns);

  return (
    <div className={styles.component}>
      {columns.map((c, index) => (
        <GqlColumnTab key={c.columnName} fragment={c} tabIndex={index} />
      ))}
    </div>
  );
}
