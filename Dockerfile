FROM public.ecr.aws/lambda/provided:al2
COPY ./build/kaytu-aws-describer ./
ENTRYPOINT [ "./kaytu-aws-describer" ]
CMD [ "./kaytu-aws-describer" ]