FROM public.ecr.aws/lambda/provided:al2
COPY ./build/og-aws-describer ./
ENTRYPOINT [ "./og-aws-describer" ]
CMD [ "./og-aws-describer" ]