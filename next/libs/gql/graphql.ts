/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from "@graphql-typed-document-node/core";
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>;
};
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>;
};
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type BackgroundImageColumn = Column & {
  __typename: "BackgroundImageColumn";
  _placeholder?: Maybe<Scalars["String"]>;
  height?: Maybe<Scalars["Int"]>;
  modal?: Maybe<Modal>;
  path?: Maybe<Scalars["String"]>;
  url?: Maybe<Scalars["String"]>;
  width?: Maybe<Scalars["Int"]>;
};

export type Column = {
  _placeholder?: Maybe<Scalars["String"]>;
};

export type ColumnVerticalPosition = "BOTTOM" | "CENTER" | "TOP";

export type ColumnWrapper = {
  __typename: "ColumnWrapper";
  column?: Maybe<Column>;
  index?: Maybe<Scalars["Int"]>;
  name?: Maybe<Scalars["String"]>;
};

export type FileHighlight = {
  __typename: "FileHighlight";
  fromLine?: Maybe<Scalars["Int"]>;
  toLine?: Maybe<Scalars["Int"]>;
};

export type FileNode = {
  __typename: "FileNode";
  filePath?: Maybe<Scalars["String"]>;
  isUpdated?: Maybe<Scalars["Boolean"]>;
  name?: Maybe<Scalars["String"]>;
  nodeType?: Maybe<FileNodeType>;
  offset?: Maybe<Scalars["Int"]>;
};

export type FileNodeType = "DIRECTORY" | "FILE";

export type ImageCentered = {
  __typename: "ImageCentered";
  height?: Maybe<Scalars["Int"]>;
  path?: Maybe<Scalars["String"]>;
  url?: Maybe<Scalars["String"]>;
  width?: Maybe<Scalars["Int"]>;
};

export type ImageDescriptionColumn = Column & {
  __typename: "ImageDescriptionColumn";
  _placeholder?: Maybe<Scalars["String"]>;
  contentsPosition?: Maybe<ColumnVerticalPosition>;
  description?: Maybe<Markdown>;
  image?: Maybe<ImageCentered>;
  order?: Maybe<ImageDescriptionOrder>;
};

export type ImageDescriptionOrder =
  | "DESCRIPTION_THEN_IMAGE"
  | "IMAGE_THEN_DESCRIPTION";

export type Markdown = {
  __typename: "Markdown";
  alignment?: Maybe<MarkdownAlignment>;
  contents?: Maybe<Scalars["String"]>;
  step?: Maybe<Scalars["String"]>;
};

export type MarkdownAlignment = "CENTER" | "LEFT";

export type MarkdownColumn = Column & {
  __typename: "MarkdownColumn";
  _placeholder?: Maybe<Scalars["String"]>;
  contentsPosition?: Maybe<ColumnVerticalPosition>;
  description?: Maybe<Markdown>;
};

export type MarkdownOld = {
  __typename: "MarkdownOld";
  contents?: Maybe<Scalars["String"]>;
  step?: Maybe<Scalars["String"]>;
};

export type Modal = {
  __typename: "Modal";
  position?: Maybe<ModalPosition>;
  text?: Maybe<Scalars["String"]>;
};

export type ModalPosition = "BOTTOM" | "CENTER" | "TOP";

export type NextAction = {
  __typename: "NextAction";
  markdown?: Maybe<MarkdownOld>;
  terminalCommand?: Maybe<TerminalCommand>;
  terminalName?: Maybe<Scalars["String"]>;
};

export type OpenFile = {
  __typename: "OpenFile";
  content?: Maybe<Scalars["String"]>;
  fileName?: Maybe<Scalars["String"]>;
  filePath?: Maybe<Scalars["String"]>;
  highlight?: Maybe<Array<Maybe<FileHighlight>>>;
  isFullContent?: Maybe<Scalars["Boolean"]>;
  language?: Maybe<Scalars["String"]>;
};

export type Page = {
  __typename: "Page";
  columns?: Maybe<Array<Maybe<ColumnWrapper>>>;
  focusColumn?: Maybe<Scalars["String"]>;
  modal?: Maybe<Modal>;
  nextStep?: Maybe<Scalars["String"]>;
  prevStep?: Maybe<Scalars["String"]>;
  step?: Maybe<Scalars["String"]>;
};

export type PageState = {
  __typename: "PageState";
  markdown?: Maybe<MarkdownOld>;
  nextAction?: Maybe<NextAction>;
  nextStep?: Maybe<Scalars["String"]>;
  prevStep?: Maybe<Scalars["String"]>;
  sourceCode?: Maybe<SourceCode>;
  step?: Maybe<Scalars["String"]>;
  terminals?: Maybe<Array<Maybe<Terminal>>>;
};

export type Query = {
  __typename: "Query";
  page?: Maybe<Page>;
  pageState?: Maybe<PageState>;
};

export type QueryPageArgs = {
  step?: InputMaybe<Scalars["String"]>;
  tutorial: Scalars["String"];
};

export type QueryPageStateArgs = {
  step?: InputMaybe<Scalars["String"]>;
};

export type SourceCode = {
  __typename: "SourceCode";
  fileTree?: Maybe<Array<Maybe<FileNode>>>;
  openFile?: Maybe<OpenFile>;
  step?: Maybe<Scalars["String"]>;
};

export type SourceCodeOpenFileArgs = {
  filePath?: InputMaybe<Scalars["String"]>;
};

export type SourceCodeColumn = Column & {
  __typename: "SourceCodeColumn";
  _placeholder?: Maybe<Scalars["String"]>;
  sourceCode?: Maybe<SourceCode>;
};

export type Terminal = {
  __typename: "Terminal";
  currentDirectory?: Maybe<Scalars["String"]>;
  name?: Maybe<Scalars["String"]>;
  nodes?: Maybe<Array<Maybe<TerminalNode>>>;
  step?: Maybe<Scalars["String"]>;
};

export type TerminalColumn = Column & {
  __typename: "TerminalColumn";
  _placeholder?: Maybe<Scalars["String"]>;
  terminal?: Maybe<Terminal>;
};

export type TerminalCommand = {
  __typename: "TerminalCommand";
  beforeExecution?: Maybe<Scalars["Boolean"]>;
  command?: Maybe<Scalars["String"]>;
};

export type TerminalElement = TerminalCommand | TerminalOutput;

export type TerminalNode = {
  __typename: "TerminalNode";
  content?: Maybe<TerminalElement>;
};

export type TerminalOutput = {
  __typename: "TerminalOutput";
  output?: Maybe<Scalars["String"]>;
};

export type BackgroundImageColumnFragmentFragment = {
  __typename: "BackgroundImageColumn";
  width?: number | null;
  height?: number | null;
  path?: string | null;
  modal?:
    | ({ __typename: "Modal"; text?: string | null } & {
        " $fragmentRefs"?: {
          ModalFrameFragmentFragment: ModalFrameFragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "BackgroundImageColumnFragmentFragment" };

export type ColumnWrapperFragmentFragment = {
  __typename: "ColumnWrapper";
  column?:
    | ({ __typename: "BackgroundImageColumn" } & {
        " $fragmentRefs"?: {
          BackgroundImageColumnFragmentFragment: BackgroundImageColumnFragmentFragment;
        };
      })
    | ({
        __typename: "ImageDescriptionColumn";
        contentsPosition?: ColumnVerticalPosition | null;
      } & {
        " $fragmentRefs"?: {
          ImageDescriptionColumnFragmentFragment: ImageDescriptionColumnFragmentFragment;
        };
      })
    | ({ __typename: "MarkdownColumn" } & {
        " $fragmentRefs"?: {
          MarkdownColumnFragmentFragment: MarkdownColumnFragmentFragment;
        };
      })
    | ({ __typename: "SourceCodeColumn" } & {
        " $fragmentRefs"?: {
          SourceCodeColumnFragmentFragment: SourceCodeColumnFragmentFragment;
        };
      })
    | ({ __typename: "TerminalColumn" } & {
        " $fragmentRefs"?: {
          TerminalColumnFragmentFragment: TerminalColumnFragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "ColumnWrapperFragmentFragment" };

export type ImageDescriptionColumnFragmentFragment = {
  __typename: "ImageDescriptionColumn";
  order?: ImageDescriptionOrder | null;
  contentsPosition?: ColumnVerticalPosition | null;
  description?:
    | ({ __typename: "Markdown" } & {
        " $fragmentRefs"?: {
          MarkdownFragmentFragment: MarkdownFragmentFragment;
        };
      })
    | null;
  image?:
    | ({ __typename: "ImageCentered" } & {
        " $fragmentRefs"?: {
          ImageCenteredFragmentFragment: ImageCenteredFragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "ImageDescriptionColumnFragmentFragment" };

export type MarkdownColumnFragmentFragment = {
  __typename: "MarkdownColumn";
  contentsPosition?: ColumnVerticalPosition | null;
  description?:
    | ({ __typename: "Markdown" } & {
        " $fragmentRefs"?: {
          MarkdownFragmentFragment: MarkdownFragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "MarkdownColumnFragmentFragment" };

export type PageColumnsFragmentFragment = {
  __typename: "Page";
  columns?: Array<
    | ({ __typename: "ColumnWrapper" } & {
        " $fragmentRefs"?: {
          ColumnWrapperFragmentFragment: ColumnWrapperFragmentFragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "PageColumnsFragmentFragment" };

export type SourceCodeColumnFragmentFragment = {
  __typename: "SourceCodeColumn";
  sourceCode?:
    | ({ __typename: "SourceCode" } & {
        " $fragmentRefs"?: {
          SourceCodeViewer_FragmentFragment: SourceCodeViewer_FragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "SourceCodeColumnFragmentFragment" };

export type TerminalColumnFragmentFragment = {
  __typename: "TerminalColumn";
  terminal?:
    | ({ __typename: "Terminal" } & {
        " $fragmentRefs"?: {
          TerminalComponent_FragmentFragment: TerminalComponent_FragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "TerminalColumnFragmentFragment" };

export type ImageCenteredFragmentFragment = {
  __typename: "ImageCentered";
  width?: number | null;
  height?: number | null;
  path?: string | null;
} & { " $fragmentName"?: "ImageCenteredFragmentFragment" };

export type MarkdownFragmentFragment = {
  __typename: "Markdown";
  contents?: string | null;
  alignment?: MarkdownAlignment | null;
} & { " $fragmentName"?: "MarkdownFragmentFragment" };

export type ModalFrameFragmentFragment = {
  __typename: "Modal";
  text?: string | null;
  position?: ModalPosition | null;
} & { " $fragmentName"?: "ModalFrameFragmentFragment" };

export type NavigationFragmentFragment = {
  __typename: "Page";
  step?: string | null;
  nextStep?: string | null;
  prevStep?: string | null;
} & { " $fragmentName"?: "NavigationFragmentFragment" };

export type SourceCodeViewer_FragmentFragment = ({
  __typename: "SourceCode";
  openFile?:
    | ({ __typename: "OpenFile" } & {
        " $fragmentRefs"?: {
          FileContentPane_FragmentFragment: FileContentPane_FragmentFragment;
        };
      })
    | null;
} & {
  " $fragmentRefs"?: {
    FileTreePane_FragmentFragment: FileTreePane_FragmentFragment;
  };
}) & { " $fragmentName"?: "SourceCodeViewer_FragmentFragment" };

export type FileNodeComponent_FragmentFragment = ({
  __typename: "FileNode";
  nodeType?: FileNodeType | null;
  name?: string | null;
  filePath?: string | null;
  offset?: number | null;
  isUpdated?: boolean | null;
} & {
  " $fragmentRefs"?: {
    FileNodeIcon_FragmentFragment: FileNodeIcon_FragmentFragment;
  };
}) & { " $fragmentName"?: "FileNodeComponent_FragmentFragment" };

export type FileNodeIcon_FragmentFragment = {
  __typename: "FileNode";
  nodeType?: FileNodeType | null;
} & { " $fragmentName"?: "FileNodeIcon_FragmentFragment" };

export type FileTreeComponent_FragmentFragment = {
  __typename: "SourceCode";
  fileTree?: Array<
    | ({ __typename: "FileNode"; filePath?: string | null } & {
        " $fragmentRefs"?: {
          FileNodeComponent_FragmentFragment: FileNodeComponent_FragmentFragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "FileTreeComponent_FragmentFragment" };

export type FileTreePane_FragmentFragment = ({ __typename: "SourceCode" } & {
  " $fragmentRefs"?: {
    FileTreeComponent_FragmentFragment: FileTreeComponent_FragmentFragment;
  };
}) & { " $fragmentName"?: "FileTreePane_FragmentFragment" };

export type FileContentPane_FragmentFragment = ({ __typename: "OpenFile" } & {
  " $fragmentRefs"?: {
    FileNameTabBar_FragmentFragment: FileNameTabBar_FragmentFragment;
    FileContentViewer_FragmentFragment: FileContentViewer_FragmentFragment;
  };
}) & { " $fragmentName"?: "FileContentPane_FragmentFragment" };

export type FileContentViewer_FragmentFragment = {
  __typename: "OpenFile";
  content?: string | null;
  language?: string | null;
  highlight?: Array<{
    __typename: "FileHighlight";
    fromLine?: number | null;
    toLine?: number | null;
  } | null> | null;
} & { " $fragmentName"?: "FileContentViewer_FragmentFragment" };

export type FileNameTab_FragmentFragment = {
  __typename: "OpenFile";
  fileName?: string | null;
} & { " $fragmentName"?: "FileNameTab_FragmentFragment" };

export type FileNameTabBar_FragmentFragment = ({ __typename: "OpenFile" } & {
  " $fragmentRefs"?: {
    FileNameTab_FragmentFragment: FileNameTab_FragmentFragment;
  };
}) & { " $fragmentName"?: "FileNameTabBar_FragmentFragment" };

export type TerminalCommand_FragmentFragment = {
  __typename: "TerminalCommand";
  command?: string | null;
  beforeExecution?: boolean | null;
} & { " $fragmentName"?: "TerminalCommand_FragmentFragment" };

export type TerminalComponent_FragmentFragment = {
  __typename: "Terminal";
  currentDirectory?: string | null;
  nodes?: Array<
    | ({ __typename: "TerminalNode" } & {
        " $fragmentRefs"?: {
          TerminalNodeComponent_FragmentFragment: TerminalNodeComponent_FragmentFragment;
        };
      })
    | null
  > | null;
} & { " $fragmentName"?: "TerminalComponent_FragmentFragment" };

export type TerminalNodeComponent_FragmentFragment = {
  __typename: "TerminalNode";
  content?:
    | ({ __typename: "TerminalCommand" } & {
        " $fragmentRefs"?: {
          TerminalCommand_FragmentFragment: TerminalCommand_FragmentFragment;
        };
      })
    | ({ __typename: "TerminalOutput" } & {
        " $fragmentRefs"?: {
          TerminalOutput_FragmentFragment: TerminalOutput_FragmentFragment;
        };
      })
    | null;
} & { " $fragmentName"?: "TerminalNodeComponent_FragmentFragment" };

export type TerminalOutput_FragmentFragment = {
  __typename: "TerminalOutput";
  output?: string | null;
} & { " $fragmentName"?: "TerminalOutput_FragmentFragment" };

export type IndexSsrPageQueryVariables = Exact<{
  tutorial: Scalars["String"];
  step?: InputMaybe<Scalars["String"]>;
  openFilePath?: InputMaybe<Scalars["String"]>;
}>;

export type IndexSsrPageQuery = {
  __typename: "Query";
  page?:
    | ({ __typename: "Page" } & {
        " $fragmentRefs"?: {
          NavigationFragmentFragment: NavigationFragmentFragment;
          PageColumnsFragmentFragment: PageColumnsFragmentFragment;
        };
      })
    | null;
};

export const MarkdownFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Markdown" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "contents" } },
          { kind: "Field", name: { kind: "Name", value: "alignment" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<MarkdownFragmentFragment, unknown>;
export const ImageCenteredFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ImageCenteredFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ImageCentered" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "width" } },
          { kind: "Field", name: { kind: "Name", value: "height" } },
          { kind: "Field", name: { kind: "Name", value: "path" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<ImageCenteredFragmentFragment, unknown>;
export const ImageDescriptionColumnFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ImageDescriptionColumnFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ImageDescriptionColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "description" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "MarkdownFragment" },
                },
              ],
            },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "image" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "ImageCenteredFragment" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "order" } },
          { kind: "Field", name: { kind: "Name", value: "contentsPosition" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<ImageDescriptionColumnFragmentFragment, unknown>;
export const ModalFrameFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ModalFrameFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Modal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "text" } },
          { kind: "Field", name: { kind: "Name", value: "position" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<ModalFrameFragmentFragment, unknown>;
export const BackgroundImageColumnFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "BackgroundImageColumnFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "BackgroundImageColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "width" } },
          { kind: "Field", name: { kind: "Name", value: "height" } },
          { kind: "Field", name: { kind: "Name", value: "path" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "modal" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "text" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "ModalFrameFragment" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<BackgroundImageColumnFragmentFragment, unknown>;
export const MarkdownColumnFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "MarkdownColumnFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "MarkdownColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "description" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "MarkdownFragment" },
                },
              ],
            },
          },
          { kind: "Field", name: { kind: "Name", value: "contentsPosition" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<MarkdownColumnFragmentFragment, unknown>;
export const TerminalCommand_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalCommand_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalCommand" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "command" } },
          { kind: "Field", name: { kind: "Name", value: "beforeExecution" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<TerminalCommand_FragmentFragment, unknown>;
export const TerminalOutput_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalOutput_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalOutput" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "output" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<TerminalOutput_FragmentFragment, unknown>;
export const TerminalNodeComponent_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "content" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalCommand" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalCommand_Fragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalOutput" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "TerminalOutput_Fragment",
                        },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<TerminalNodeComponent_FragmentFragment, unknown>;
export const TerminalComponent_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Terminal" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "currentDirectory" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "nodes" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: {
                    kind: "Name",
                    value: "TerminalNodeComponent_Fragment",
                  },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<TerminalComponent_FragmentFragment, unknown>;
export const TerminalColumnFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "TerminalColumnFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "TerminalColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "terminal" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "TerminalComponent_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<TerminalColumnFragmentFragment, unknown>;
export const FileNodeIcon_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNodeIcon_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<FileNodeIcon_FragmentFragment, unknown>;
export const FileNodeComponent_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNodeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "FileNode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileNodeIcon_Fragment" },
          },
          { kind: "Field", name: { kind: "Name", value: "nodeType" } },
          { kind: "Field", name: { kind: "Name", value: "name" } },
          { kind: "Field", name: { kind: "Name", value: "filePath" } },
          { kind: "Field", name: { kind: "Name", value: "offset" } },
          { kind: "Field", name: { kind: "Name", value: "isUpdated" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<FileNodeComponent_FragmentFragment, unknown>;
export const FileTreeComponent_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileTreeComponent_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "fileTree" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "filePath" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "FileNodeComponent_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<FileTreeComponent_FragmentFragment, unknown>;
export const FileTreePane_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileTreePane_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileTreeComponent_Fragment" },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<FileTreePane_FragmentFragment, unknown>;
export const FileNameTab_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNameTab_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "fileName" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<FileNameTab_FragmentFragment, unknown>;
export const FileNameTabBar_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileNameTabBar_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileNameTab_Fragment" },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<FileNameTabBar_FragmentFragment, unknown>;
export const FileContentViewer_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileContentViewer_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "content" } },
          { kind: "Field", name: { kind: "Name", value: "language" } },
          {
            kind: "Field",
            name: { kind: "Name", value: "highlight" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "fromLine" } },
                { kind: "Field", name: { kind: "Name", value: "toLine" } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<FileContentViewer_FragmentFragment, unknown>;
export const FileContentPane_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "FileContentPane_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "OpenFile" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileNameTabBar_Fragment" },
          },
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileContentViewer_Fragment" },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<FileContentPane_FragmentFragment, unknown>;
export const SourceCodeViewer_FragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "SourceCodeViewer_Fragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCode" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "FragmentSpread",
            name: { kind: "Name", value: "FileTreePane_Fragment" },
          },
          {
            kind: "Field",
            name: { kind: "Name", value: "openFile" },
            arguments: [
              {
                kind: "Argument",
                name: { kind: "Name", value: "filePath" },
                value: {
                  kind: "Variable",
                  name: { kind: "Name", value: "openFilePath" },
                },
              },
            ],
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "FileContentPane_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<SourceCodeViewer_FragmentFragment, unknown>;
export const SourceCodeColumnFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "SourceCodeColumnFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "SourceCodeColumn" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "sourceCode" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "SourceCodeViewer_Fragment" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<SourceCodeColumnFragmentFragment, unknown>;
export const ColumnWrapperFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "ColumnWrapperFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "ColumnWrapper" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "column" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "ImageDescriptionColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "ImageDescriptionColumnFragment",
                        },
                      },
                      {
                        kind: "Field",
                        name: { kind: "Name", value: "contentsPosition" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "BackgroundImageColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "BackgroundImageColumnFragment",
                        },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "MarkdownColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "MarkdownColumnFragment" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "TerminalColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: { kind: "Name", value: "TerminalColumnFragment" },
                      },
                    ],
                  },
                },
                {
                  kind: "InlineFragment",
                  typeCondition: {
                    kind: "NamedType",
                    name: { kind: "Name", value: "SourceCodeColumn" },
                  },
                  selectionSet: {
                    kind: "SelectionSet",
                    selections: [
                      {
                        kind: "FragmentSpread",
                        name: {
                          kind: "Name",
                          value: "SourceCodeColumnFragment",
                        },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<ColumnWrapperFragmentFragment, unknown>;
export const PageColumnsFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "PageColumnsFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "columns" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "ColumnWrapperFragment" },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<PageColumnsFragmentFragment, unknown>;
export const NavigationFragmentFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "NavigationFragment" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "Page" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "step" } },
          { kind: "Field", name: { kind: "Name", value: "nextStep" } },
          { kind: "Field", name: { kind: "Name", value: "prevStep" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<NavigationFragmentFragment, unknown>;
export const IndexSsrPageDocument = {
  kind: "Document",
  definitions: [
    {
      kind: "OperationDefinition",
      operation: "query",
      name: { kind: "Name", value: "IndexSsrPage" },
      variableDefinitions: [
        {
          kind: "VariableDefinition",
          variable: {
            kind: "Variable",
            name: { kind: "Name", value: "tutorial" },
          },
          type: {
            kind: "NonNullType",
            type: {
              kind: "NamedType",
              name: { kind: "Name", value: "String" },
            },
          },
        },
        {
          kind: "VariableDefinition",
          variable: { kind: "Variable", name: { kind: "Name", value: "step" } },
          type: { kind: "NamedType", name: { kind: "Name", value: "String" } },
        },
        {
          kind: "VariableDefinition",
          variable: {
            kind: "Variable",
            name: { kind: "Name", value: "openFilePath" },
          },
          type: { kind: "NamedType", name: { kind: "Name", value: "String" } },
        },
      ],
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "page" },
            arguments: [
              {
                kind: "Argument",
                name: { kind: "Name", value: "tutorial" },
                value: {
                  kind: "Variable",
                  name: { kind: "Name", value: "tutorial" },
                },
              },
              {
                kind: "Argument",
                name: { kind: "Name", value: "step" },
                value: {
                  kind: "Variable",
                  name: { kind: "Name", value: "step" },
                },
              },
            ],
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                { kind: "Field", name: { kind: "Name", value: "__typename" } },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "NavigationFragment" },
                },
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "PageColumnsFragment" },
                },
              ],
            },
          },
        ],
      },
    },
    ...NavigationFragmentFragmentDoc.definitions,
    ...PageColumnsFragmentFragmentDoc.definitions,
    ...ColumnWrapperFragmentFragmentDoc.definitions,
    ...ImageDescriptionColumnFragmentFragmentDoc.definitions,
    ...MarkdownFragmentFragmentDoc.definitions,
    ...ImageCenteredFragmentFragmentDoc.definitions,
    ...BackgroundImageColumnFragmentFragmentDoc.definitions,
    ...ModalFrameFragmentFragmentDoc.definitions,
    ...MarkdownColumnFragmentFragmentDoc.definitions,
    ...TerminalColumnFragmentFragmentDoc.definitions,
    ...TerminalComponent_FragmentFragmentDoc.definitions,
    ...TerminalNodeComponent_FragmentFragmentDoc.definitions,
    ...TerminalCommand_FragmentFragmentDoc.definitions,
    ...TerminalOutput_FragmentFragmentDoc.definitions,
    ...SourceCodeColumnFragmentFragmentDoc.definitions,
    ...SourceCodeViewer_FragmentFragmentDoc.definitions,
    ...FileTreePane_FragmentFragmentDoc.definitions,
    ...FileTreeComponent_FragmentFragmentDoc.definitions,
    ...FileNodeComponent_FragmentFragmentDoc.definitions,
    ...FileNodeIcon_FragmentFragmentDoc.definitions,
    ...FileContentPane_FragmentFragmentDoc.definitions,
    ...FileNameTabBar_FragmentFragmentDoc.definitions,
    ...FileNameTab_FragmentFragmentDoc.definitions,
    ...FileContentViewer_FragmentFragmentDoc.definitions,
  ],
} as unknown as DocumentNode<IndexSsrPageQuery, IndexSsrPageQueryVariables>;
