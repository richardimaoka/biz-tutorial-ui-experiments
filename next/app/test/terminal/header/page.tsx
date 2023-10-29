import { TerminalHeaderGql } from "@/app/components/terminal2/header/TerminalHeaderGql";
import { TerminalHeaderGqlFragment } from "@/libs/gql/graphql";

export default function Page() {
  const fragmentt: TerminalHeaderGqlFragment = {
    __typename: "Terminal2",
    currentDirectory: "dirdir",
  };

  type a = {
    " $fragmentRefs"?:
      | {
          TerminalHeaderGqlFragment: TerminalHeaderGqlFragment;
        }
      | undefined;
  };

  function typeCast(f: TerminalHeaderGqlFragment): a {
    return {
      " $fragmentRefs": {
        TerminalHeaderGqlFragment: f,
      },
      ...f,
    };
  }

  console.log(typeCast(fragmentt));

  return (
    <div>
      <TerminalHeaderGql
        fragment={typeCast(fragmentt)}
        // fragment={{
        //   " $fragmentRefs": {
        //     TerminalHeaderGqlFragment: fragmentt,
        //   },
        //   ...fragmentt,
        // }}
      />
    </div>
  );
}
