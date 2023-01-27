import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { TerminalCommandComponent } from "./TerminalCommandComponent";
import { TerminalCommandWritingComponent } from "./TerminalCommandWritingComponent";
import { TerminalOutputComponent } from "./TerminalOutputComponent";

const TerminalComponent_Fragment = graphql(`
  fragment TerminalComponent_Fragment on Terminal {
    currentDirectory
    elements {
      __typename
      ... on TerminalCommand {
        ...TerminalCommand_Fragment
        ...TerminalCommandWriting_Fragment
      }
      ... on TerminalOutput {
        ...TerminalOutput_Fragment
      }
    }
  }
`);

export interface TerminalComponentProps {
  fragment: FragmentType<typeof TerminalComponent_Fragment>;
}

export const TerminalComponent = (
  props: TerminalComponentProps
): JSX.Element => {
  const fragment = useFragment(TerminalComponent_Fragment, props.fragment);

  return (
    <>
      {fragment.elements?.map((elem, index) => {
        if (elem) {
          switch (elem.__typename) {
            case "TerminalCommand":
              return fragment.elements?.length === index + 1 ? (
                //last element
                <TerminalCommandWritingComponent key={index} fragment={elem} />
              ) : (
                <TerminalCommandComponent key={index} fragment={elem} />
              );
            case "TerminalOutput":
              return <TerminalOutputComponent key={index} fragment={elem} />;
            case "TerminalCommandSet":
              return <></>;
          }
        } else {
          return <></>;
        }
      })}
    </>
  );
};
