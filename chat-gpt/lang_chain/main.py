from langchain_core.output_parsers import StrOutputParser
from langchain_core.prompts import ChatPromptTemplate
from langchain_openai import ChatOpenAI

llm = ChatOpenAI(model_name="gpt-3.5-turbo", temperature=0)

template = """
次の文章に誤字がないか調べて。誤字があれば訂正してください。
{sentences_before_check}
"""

prompt = ChatPromptTemplate.from_messages(
    [
       ("system", "あなたは優秀な校正者です。"),
       ("user", template)
    ]
)

output_parser = StrOutputParser()

chain = prompt | llm | output_parser

print(chain.invoke({
    "sentences_before_check": "こんんにちわ、よっぴです。",
}))