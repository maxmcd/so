for x in range(1,10):
    with open("thing.txt","a") as gene_list:
        gene_list.write("\n subj3: {}".format(x))
        print(x)
        gene_list.close()
