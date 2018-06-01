import os, shutil
from distutils.dir_util import copy_tree

scriptsDir = os.path.dirname(os.path.realpath(__file__))
#print("Scripts folder: " + scriptsDir)
srcDir = scriptsDir + "/../src"
commonDir = srcDir + "/common"
targetDir = srcDir + "/myproject/common"
shutil.rmtree(targetDir, True)
print("About to copy folder '" + commonDir + "' to '" + targetDir + "'")

if os.path.isdir(commonDir) and os.path.exists(commonDir):
    copy_tree(commonDir, targetDir)
    print("Completed copying common folder successfully!")
else:
    print("Something is wrong... the common folder does not exist!")
