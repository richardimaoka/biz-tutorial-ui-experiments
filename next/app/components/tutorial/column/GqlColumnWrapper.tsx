import { GqlTerminalColumn } from "../../terminal2/GqlTerminalColumn";
import styles from "./GqlColumnWrapper.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment GqlColumnWrapper on ColumnWrapper2 {
    columnName
    column {
      # if you forget this, the resulting fragment will have __typename = undefined
      __typename
      #
      # for each column type
      #
      ... on TerminalColumn2 {
        ...GqlTerminalColumn
      }
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlColumnWrapper(props: Props): JSX.Element {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const column = fragment.column;

  if (!column.__typename) {
    throw new Error(
      "__typename got undefined - define __typename in GraphQL fragment/query"
    );
  }

  switch (column.__typename) {
    case "TerminalColumn2":
      return (
        <div className={styles.component}>
          <GqlTerminalColumn fragment={column} selectIndex={0} />
        </div>
      );
  }
}
