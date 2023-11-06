import { nonNullArray } from "@/libs/nonNullArray";
import styles from "./GqlColumnWrappers.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { GqlColumnWrapper } from "./GqlColumnWrapper";
import { Carousel } from "../../carousel/Carousel";
import { columnWidthPx } from "../definitions";

const fragmentDefinition = graphql(`
  fragment GqlColumnWrappers on Page2 {
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
    return <></>;
  }

  const columns = nonNullArray(fragment.columns);

  return (
    <Carousel currentIndex={0} columnWidth={columnWidthPx}>
      <div className={styles.component}>
        {columns.map((c) => (
          <div key={c.columnName} className={styles.column}>
            <GqlColumnWrapper fragment={c} />
          </div>
        ))}
      </div>
    </Carousel>
  );
}
