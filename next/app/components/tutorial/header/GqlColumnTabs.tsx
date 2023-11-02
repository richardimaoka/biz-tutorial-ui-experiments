import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { nonNullArray } from "@/libs/nonNullArray";
import { GqlColumnTab } from "./GqlColumnTab";
import styles from "./GqlTutorialHeader.module.css";

const fragmentDefinition = graphql(`
  fragment GqlColumnTabs on Page2 {
    columns {
      columnName
      ...GqlColumnTab
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  selectTab?: string;
}

export function GqlColumnTabs(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.columns) {
    return <></>;
  }

  const columns = nonNullArray(fragment.columns);

  return (
    <div className={styles.component}>
      {columns.map((c) => (
        <GqlColumnTab key={c.columnName} fragment={c} />
      ))}
    </div>
  );
}
