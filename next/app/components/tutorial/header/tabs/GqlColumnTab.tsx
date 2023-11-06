import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlColumnTab.module.css";
import { GqlColumnTabIcon } from "./GqlColumnTabIcon";
import { LinkSearchParams } from "@/app/components/link/LinkSearchParams";

const fragmentDefinition = graphql(`
  fragment GqlColumnTab on ColumnWrapper2 {
    columnName
    columnDisplayName
    ...GqlColumnTabIcon
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  isSelected?: boolean;
}

export function GqlColumnTab(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  const selectStyle = props.isSelected ? styles.selected : styles.unselected;
  const outerClassName = `${styles.component} ${selectStyle}`;

  const searchParams = fragment.columnName
    ? {
        column: fragment.columnName,
      }
    : ({} as Record<string, string>);

  return (
    <LinkSearchParams searchParams={searchParams}>
      <div className={outerClassName}>
        <span className={styles.smartphone}>
          <GqlColumnTabIcon fragment={fragment} />
        </span>
        <span className={styles.desktop}>{fragment.columnName}</span>
      </div>
    </LinkSearchParams>
  );
}
