import { TerminalHeaderGql } from "@/app/components/terminal2/header/TerminalHeaderGql";
import { TerminalHeaderGqlFragment } from "@/libs/gql/graphql";
import { TypedDocumentNode } from "@graphql-typed-document-node/core";

import { FragmentType } from "@/libs/gql";

type b = FragmentType<TypedDocumentNode<TerminalHeaderGqlFragment, unknown>>;

function typeCast(f: TerminalHeaderGqlFragment): b {
  return {
    " $fragmentRefs": {
      TerminalHeaderGqlFragment: f,
    },
    ...f,
  };
}

export default function Page() {
  type TT = TypedDocumentNode<TerminalHeaderGqlFragment, unknown>;

  const fragmentt: TerminalHeaderGqlFragment = {
    __typename: "Terminal2",
    currentDirectory: "dirdir",
  };

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
