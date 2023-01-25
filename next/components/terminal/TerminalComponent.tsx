import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { TerminalElementComponent } from "./TerminalElementComponent";

const TerminalComponent_Fragment = graphql(`
  fragment TerminalComponent_Fragment on Terminal {
    currentDirectory
    elements {
      ...TerminalElementComponent_Fragment
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
      {fragment.elements?.map(
        (elem, index) =>
          elem && <TerminalElementComponent key={index} fragment={elem} />
      )}
    </>
  );
};
