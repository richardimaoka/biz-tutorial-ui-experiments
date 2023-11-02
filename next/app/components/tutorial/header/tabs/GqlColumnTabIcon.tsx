import { TerminalIcon } from "../../../icons/TerminalIcon";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment GqlColumnTabIcon on ColumnWrapper2 {
    column {
      __typename
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlColumnTabIcon(props: Props): JSX.Element {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  switch (fragment.column.__typename) {
    case "TerminalColumn2":
      return <TerminalIcon />;
  }
}
