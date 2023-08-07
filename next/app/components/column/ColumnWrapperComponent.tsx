import { FragmentType, graphql, useFragment } from "@/libs/gql";

import styles from "./style.module.css";
import { SourceCodeColumn } from "../sourcecode/SourceCodeColumn";
import { TerminalColumn } from "../terminal/TerminalColumn";

const fragmentDefinition = graphql(`
  fragment ColumnWrapperComponent_Fragment on ColumnWrapper {
    name
    column {
      __typename
      ... on SourceCodeColumn {
        ...SourceCodeColumn_Fragment
      }

      ... on TerminalColumn {
        ...TerminalColumn_Fragment
      }
    }
  }
`);

interface ColumnWrapperComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  skipAnimation?: boolean;
}

export const ColumnWrapperComponent = (props: ColumnWrapperComponentProps) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  const Column = (): JSX.Element => {
    if (!fragment?.column?.__typename) {
      return <div></div>;
    }

    switch (fragment.column.__typename) {
      case "SourceCodeColumn":
        return <SourceCodeColumn fragment={fragment.column} />;
      case "TerminalColumn":
        return (
          <TerminalColumn
            fragment={fragment.column}
            skipAnimation={props.skipAnimation}
          />
        );
      default:
        return <div>default</div>;
    }
  };

  return <Column />;
};
